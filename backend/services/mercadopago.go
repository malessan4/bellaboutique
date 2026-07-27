package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"bellaboutique/models"
)

type MPItem struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	CurrencyID  string  `json:"currency_id"`
}

type MPPayer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MPBackURLs struct {
	Success string `json:"success"`
	Failure string `json:"failure"`
	Pending string `json:"pending"`
}

type MPPreferenceRequest struct {
	Items             []MPItem   `json:"items"`
	Payer             MPPayer    `json:"payer,omitempty"`
	BackURLs          MPBackURLs `json:"back_urls,omitempty"`
	AutoReturn        string     `json:"auto_return,omitempty"`
	ExternalReference string     `json:"external_reference"`
	NotificationURL   string     `json:"notification_url,omitempty"`
}

type MPPreferenceResponse struct {
	ID               string `json:"id"`
	InitPoint        string `json:"init_point"`
	SandboxInitPoint string `json:"sandbox_init_point"`
}

type MPPaymentResponse struct {
	ID                int64  `json:"id"`
	Status            string `json:"status"`
	StatusDetail      string `json:"status_detail"`
	ExternalReference string `json:"external_reference"`
}

func CreateMPPreference(order *models.Order, items []models.OrderItem, accessToken, frontendURL string) (*MPPreferenceResponse, error) {
	mpItems := make([]MPItem, 0, len(items)+1)
	for _, item := range items {
		mpItems = append(mpItems, MPItem{
			ID:          fmt.Sprintf("%d", item.ProductID),
			Title:       item.ProductName,
			Description: fmt.Sprintf("Talle: %s | Color: %s", item.Size, item.Color),
			Quantity:    item.Quantity,
			UnitPrice:   item.Price,
			CurrencyID:  "ARS",
		})
	}
	if order.ShippingCost > 0 {
		mpItems = append(mpItems, MPItem{
			ID: "envio", Title: "Costo de Envio",
			Quantity: 1, UnitPrice: order.ShippingCost, CurrencyID: "ARS",
		})
	}

	orderRef := fmt.Sprintf("%d", order.ID)
	prefReq := MPPreferenceRequest{
		Items:  mpItems,
		ExternalReference: orderRef,
	}

	// back_urls y auto_return solo funcionan con URLs publicas (no localhost)
	if !strings.Contains(frontendURL, "localhost") {
		prefReq.BackURLs = MPBackURLs{
			Success: frontendURL + "/pago/exitoso?order=" + orderRef,
			Failure: frontendURL + "/pago/fallido?order=" + orderRef,
			Pending: frontendURL + "/pago/pendiente?order=" + orderRef,
		}
		prefReq.AutoReturn = "approved"
		// Le decimos a MP donde avisarnos que el pago fue aprobado (Webhook)
		prefReq.NotificationURL = "https://bellaboutique.onrender.com/api/payments/webhook"
	}

	body, err := json.Marshal(prefReq)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://api.mercadopago.com/checkout/preferences", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error MercadoPago (%d): %s", resp.StatusCode, string(respBody))
	}

	var prefResp MPPreferenceResponse
	if err := json.Unmarshal(respBody, &prefResp); err != nil {
		return nil, err
	}
	return &prefResp, nil
}

func GetMPPayment(paymentID, accessToken string) (*MPPaymentResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.mercadopago.com/v1/payments/%s", paymentID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var payment MPPaymentResponse
	json.NewDecoder(resp.Body).Decode(&payment)
	return &payment, nil
}
