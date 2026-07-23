package database

import (
	"log"

	"bellaboutique/config"
	"bellaboutique/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB, cfg *config.Config) {
	var count int64
	db.Model(&models.Category{}).Count(&count)
	if count > 0 {
		log.Println("Base de datos ya tiene datos, omitiendo seed")
		return
	}
	log.Println("Sembrando datos iniciales...")

	categories := []models.Category{
		{Name: "Lenceria", Slug: "lenceria", Description: "Sujetadores, bragas, bodys y corses de la mas alta calidad", Image: "https://images.unsplash.com/photo-1617038260897-41a1f14a8ca0?w=600&q=80"},
		{Name: "Ropa Casual", Slug: "ropa-casual", Description: "Moda femenina para el dia a dia con estilo y comodidad", Image: "https://images.unsplash.com/photo-1469334031218-e382a71b716b?w=600&q=80"},
		{Name: "Pijamas y Loungewear", Slug: "pijamas-loungewear", Description: "Conjuntos comodos para descansar en casa con elegancia", Image: "https://images.unsplash.com/photo-1545291730-faff8ca1d4b0?w=600&q=80"},
		{Name: "Accesorios", Slug: "accesorios", Description: "Complementos perfectos para realzar cualquier look", Image: "https://images.unsplash.com/photo-1584917865442-de89df76afd3?w=600&q=80"},
	}
	for i := range categories {
		db.Create(&categories[i])
	}

	sale1 := 6800.0
	sale2 := 11000.0
	sale3 := 10000.0
	sale4 := 16000.0

	products := []models.Product{
		// ---- LENCERIA (CategoryID: 1) ----
		{
			Name: "Soutien con Encaje Negro", Slug: "soutien-encaje-negro",
			Description: "Soutien de encaje premium con aro y copa moldeada. Confeccionado en 80% nylon y 20% elastano. Ajuste perfecto para todo el dia.",
			Price: 9500, SalePrice: &sale1,
			Images:     []string{"https://images.unsplash.com/photo-1617038260897-41a1f14a8ca0?w=600&q=80"},
			CategoryID: 1,
			Sizes:      []string{"32A", "32B", "34A", "34B", "34C", "36B", "36C"},
			Colors:     []string{"Negro", "Nude", "Bordo"},
			Stock: 50, Featured: true,
			Tags: []string{"encaje", "push-up", "soutien"},
		},
		{
			Name: "Conjunto de Lenceria Floral", Slug: "conjunto-lenceria-floral",
			Description: "Set completo con soutien y bombacha bordados con motivo floral artesanal. Romantico y sofisticado. Ideal para regalar.",
			Price: 15500,
			Images:     []string{"https://images.unsplash.com/photo-1558618666-fcd25c85cd64?w=600&q=80"},
			CategoryID: 1,
			Sizes:      []string{"XS", "S", "M", "L"},
			Colors:     []string{"Rosa", "Marfil", "Negro"},
			Stock: 35, Featured: true,
			Tags: []string{"conjunto", "floral", "romantico"},
		},
		{
			Name: "Body de Seda con Tiras", Slug: "body-seda-tiras",
			Description: "Body en seda natural con tiras finas ajustables. Elegante para usar como prenda exterior o interior. Cierre a presion en la base.",
			Price: 18000,
			Images:     []string{"https://images.unsplash.com/photo-1566479179817-1af7a0f1ab76?w=600&q=80"},
			CategoryID: 1,
			Sizes:      []string{"XS", "S", "M", "L", "XL"},
			Colors:     []string{"Blanco", "Negro", "Champagne", "Rosa"},
			Stock: 25, Featured: false,
			Tags: []string{"body", "seda", "elegante"},
		},
		{
			Name: "Corse Vintage Bordado", Slug: "corse-vintage-bordado",
			Description: "Corse con varillas y bordado floral vintage. Cierre de ganchos en la parte trasera. Look artistico y sofisticado para ocasiones especiales.",
			Price: 24000,
			Images:     []string{"https://images.unsplash.com/photo-1594938298603-c8148c4b4def?w=600&q=80"},
			CategoryID: 1,
			Sizes:      []string{"XS", "S", "M", "L"},
			Colors:     []string{"Negro", "Bordo", "Azul Noche"},
			Stock: 20, Featured: true,
			Tags: []string{"corse", "vintage", "premium"},
		},
		{
			Name: "Bombacha de Encaje Francesa", Slug: "bombacha-encaje-francesa",
			Description: "Bombacha tipo francesa en encaje de Calais. Maximo confort con elegancia clasica. Tiro bajo, corte ajustado.",
			Price: 5500,
			Images:     []string{"https://images.unsplash.com/photo-1612336307429-8a898d10e223?w=600&q=80"},
			CategoryID: 1,
			Sizes:      []string{"XS", "S", "M", "L", "XL"},
			Colors:     []string{"Negro", "Nude", "Rosa", "Bordo"},
			Stock: 80, Featured: false,
			Tags: []string{"bombacha", "encaje", "francesa"},
		},
		// ---- ROPA CASUAL (CategoryID: 2) ----
		{
			Name: "Vestido Midi Floral", Slug: "vestido-midi-floral",
			Description: "Vestido midi con estampado floral primaveral. Escote en V, mangas abullonadas. Perfecto para looks casuales y eventos al aire libre.",
			Price: 19500, SalePrice: &sale2,
			Images:     []string{"https://images.unsplash.com/photo-1515372039744-b8f02a3ae446?w=600&q=80"},
			CategoryID: 2,
			Sizes:      []string{"XS", "S", "M", "L", "XL"},
			Colors:     []string{"Floral Rosa", "Floral Azul", "Floral Neutro"},
			Stock: 40, Featured: true,
			Tags: []string{"vestido", "floral", "midi"},
		},
		{
			Name: "Blusa Off-Shoulder de Seda", Slug: "blusa-off-shoulder-seda",
			Description: "Blusa off-shoulder en seda lavable. Diseno asimetrico moderno. Versatil y elegante, ideal para multiples ocasiones.",
			Price: 12000,
			Images:     []string{"https://images.unsplash.com/photo-1434389677669-e08b4cac3105?w=600&q=80"},
			CategoryID: 2,
			Sizes:      []string{"XS", "S", "M", "L"},
			Colors:     []string{"Blanco", "Negro", "Rosa Palo", "Celeste"},
			Stock: 45, Featured: false,
			Tags: []string{"blusa", "off-shoulder", "seda"},
		},
		{
			Name: "Pantalon Wide Leg Negro", Slug: "pantalon-wide-leg-negro",
			Description: "Pantalon de pierna ancha en tela fluida. Cintura alta elastizada. El infaltable de temporada. Silueta estilizada y moderna.",
			Price: 16500,
			Images:     []string{"https://images.unsplash.com/photo-1506629082955-511b1aa562c8?w=600&q=80"},
			CategoryID: 2,
			Sizes:      []string{"XS", "S", "M", "L", "XL", "XXL"},
			Colors:     []string{"Negro", "Beige", "Blanco"},
			Stock: 55, Featured: true,
			Tags: []string{"pantalon", "wide-leg", "tendencia"},
		},
		{
			Name: "Conjunto Co-ord Beige", Slug: "conjunto-coord-beige",
			Description: "Set de dos piezas: top cropped y pantalon a tono en tela premium. Perfecto para multiples ocasiones, de dia o de noche.",
			Price: 28000,
			Images:     []string{"https://images.unsplash.com/photo-1490481651871-ab68de25d43d?w=600&q=80"},
			CategoryID: 2,
			Sizes:      []string{"XS", "S", "M", "L"},
			Colors:     []string{"Beige", "Blanco", "Negro"},
			Stock: 30, Featured: true,
			Tags: []string{"conjunto", "coord", "minimalista"},
		},
		{
			Name: "Cardigan Oversize Camel", Slug: "cardigan-oversize-camel",
			Description: "Cardigan oversize en punto suave. Ideal para dias frescos. Largo con bolsillos, tejido de alta calidad que no pica.",
			Price: 14500,
			Images:     []string{"https://images.unsplash.com/photo-1591047139829-d91aecb6caea?w=600&q=80"},
			CategoryID: 2,
			Sizes:      []string{"S", "M", "L", "XL"},
			Colors:     []string{"Camel", "Gris", "Crema", "Negro"},
			Stock: 60, Featured: false,
			Tags: []string{"cardigan", "oversize", "punto"},
		},
		// ---- PIJAMAS Y LOUNGEWEAR (CategoryID: 3) ----
		{
			Name: "Pijama de Saten Rosa", Slug: "pijama-saten-rosa",
			Description: "Conjunto de pijama en saten suave. Pantalon largo y camisa de manga corta con botones nacarados y bolsillo pecho.",
			Price: 18000, SalePrice: &sale3,
			Images:     []string{"https://images.unsplash.com/photo-1545291730-faff8ca1d4b0?w=600&q=80"},
			CategoryID: 3,
			Sizes:      []string{"XS", "S", "M", "L", "XL"},
			Colors:     []string{"Rosa", "Malva", "Azul Cielo", "Blanco"},
			Stock: 45, Featured: true,
			Tags: []string{"pijama", "saten", "comodo"},
		},
		{
			Name: "Conjunto Loungewear Algodon", Slug: "conjunto-loungewear-algodon",
			Description: "Set de buzo y pantalon en french terry suave. Perfecto para relajarse en casa con estilo. Lavable en lavarropas.",
			Price: 13500,
			Images:     []string{"https://images.unsplash.com/photo-1564257631407-4deb1f99d992?w=600&q=80"},
			CategoryID: 3,
			Sizes:      []string{"XS", "S", "M", "L", "XL", "XXL"},
			Colors:     []string{"Gris Melange", "Rosa Palo", "Verde Salvia", "Beige"},
			Stock: 70, Featured: false,
			Tags: []string{"loungewear", "algodon", "comodo"},
		},
		{
			Name: "Camison de Seda Marfil", Slug: "camison-seda-marfil",
			Description: "Camison largo en seda natural color marfil. Detalle de encaje en escote y bajo. Romantico y elegante para el descanso.",
			Price: 22000,
			Images:     []string{"https://images.unsplash.com/photo-1568252542512-9fe8fe9c87bb?w=600&q=80"},
			CategoryID: 3,
			Sizes:      []string{"XS", "S", "M", "L"},
			Colors:     []string{"Marfil", "Rosa Champagne", "Blanco"},
			Stock: 25, Featured: true,
			Tags: []string{"camison", "seda", "romantico"},
		},
		{
			Name: "Kimono de Seda Floral", Slug: "kimono-seda-floral",
			Description: "Kimono en seda estampada con motivos florales japoneses. Manga larga, largo midi. Luxe y atemporal, ideal como ropa de casa elegante.",
			Price: 26000,
			Images:     []string{"https://images.unsplash.com/photo-1552664688-cf412ec27db2?w=600&q=80"},
			CategoryID: 3,
			Sizes:      []string{"S/M", "L/XL"},
			Colors:     []string{"Floral Rosa", "Floral Negro", "Floral Azul"},
			Stock: 20, Featured: true,
			Tags: []string{"kimono", "seda", "floral"},
		},
		{
			Name: "Pantuflas Bombon Borreguito", Slug: "pantuflas-bombon-borreguito",
			Description: "Pantuflas tipo borreguito con bombon decorativo. Suela antideslizante. Abrigadas y adorables. El regalo perfecto.",
			Price: 7500,
			Images:     []string{"https://images.unsplash.com/photo-1512418490979-92798cec1380?w=600&q=80"},
			CategoryID: 3,
			Sizes:      []string{"35-36", "37-38", "39-40", "41-42"},
			Colors:     []string{"Rosa", "Blanco", "Beige", "Lila"},
			Stock: 90, Featured: false,
			Tags: []string{"pantuflas", "borreguito", "regalo"},
		},
		// ---- ACCESORIOS (CategoryID: 4) ----
		{
			Name: "Cartera Tote de Cuero Blanco", Slug: "cartera-tote-cuero-blanco",
			Description: "Cartera tote en cuero legitimo blanco. Gran capacidad, cierre magnetico, asa superior y bandolera removible.",
			Price: 32000, SalePrice: &sale4,
			Images:     []string{"https://images.unsplash.com/photo-1584917865442-de89df76afd3?w=600&q=80"},
			CategoryID: 4,
			Sizes:      []string{"Unico"},
			Colors:     []string{"Blanco", "Beige", "Negro"},
			Stock: 20, Featured: true,
			Tags: []string{"cartera", "cuero", "tote"},
		},
		{
			Name: "Panuelo de Seda Estampado", Slug: "panuelo-seda-estampado",
			Description: "Panuelo 100% seda con estampado exclusivo Bella Boutique. Multiuso: cuello, cabello o bolsillo. 60x60cm.",
			Price: 8500,
			Images:     []string{"https://images.unsplash.com/photo-1601924994987-69e26d50dc26?w=600&q=80"},
			CategoryID: 4,
			Sizes:      []string{"60x60cm"},
			Colors:     []string{"Floral Rosa", "Geometrico Azul", "Animal Print"},
			Stock: 50, Featured: false,
			Tags: []string{"panuelo", "seda", "accesorio"},
		},
		{
			Name: "Collar Delicado Rose Gold", Slug: "collar-delicado-rose-gold",
			Description: "Collar delicado en acero quirurgico con bano de oro 18k. Cadena fina con dije de estrella o corazon. Hipoalergenico.",
			Price: 9500,
			Images:     []string{"https://images.unsplash.com/photo-1599643478518-a784e5dc4c8f?w=600&q=80"},
			CategoryID: 4,
			Sizes:      []string{"40cm", "45cm", "50cm"},
			Colors:     []string{"Dorado", "Plateado", "Rose Gold"},
			Stock: 60, Featured: true,
			Tags: []string{"collar", "joyeria", "delicado"},
		},
		{
			Name: "Anteojos de Sol Cat Eye", Slug: "anteojos-cat-eye",
			Description: "Gafas de sol estilo cat eye con montura acetato y lentes polarizadas UV400. Un icono de la moda femenina.",
			Price: 14500,
			Images:     []string{"https://images.unsplash.com/photo-1583743814966-8936f5b7be1a?w=600&q=80"},
			CategoryID: 4,
			Sizes:      []string{"Unico"},
			Colors:     []string{"Negro", "Tortoise", "Rosa"},
			Stock: 35, Featured: false,
			Tags: []string{"anteojos", "cat-eye", "polarizado"},
		},
		{
			Name: "Perfume Floral Oriental 50ml", Slug: "perfume-floral-oriental",
			Description: "Fragancia artesanal con notas de rosa, jazmin y ambar. Concentracion EDP. Duracion: 8-10 horas. Frasco coleccionable.",
			Price: 35000,
			Images:     []string{"https://images.unsplash.com/photo-1541643600914-78b084683702?w=600&q=80"},
			CategoryID: 4,
			Sizes:      []string{"30ml", "50ml", "100ml"},
			Colors:     []string{"Unico"},
			Stock: 25, Featured: true,
			Tags: []string{"perfume", "floral", "oriental"},
		},
	}
	for i := range products {
		db.Create(&products[i])
	}

	// Admin user
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(cfg.AdminPassword), bcrypt.DefaultCost)
	admin := models.User{
		Email:    cfg.AdminEmail,
		Password: string(hashedPwd),
		Name:     "Administrador",
		Role:     "admin",
		IsActive: true,
	}
	db.Where(models.User{Email: cfg.AdminEmail}).FirstOrCreate(&admin)
	log.Println("Datos iniciales sembrados exitosamente!")
}
