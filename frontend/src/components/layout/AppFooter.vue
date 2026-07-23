<template>
  <footer class="footer">
    <div class="container">
      <div class="footer__grid">
        <div>
          <div class="footer__brand">Bella Boutique</div>
          <p class="footer__tag">Moda femenina & lencería premium para la mujer moderna.</p>
          <div class="footer__social">
            <a href="#">Instagram</a>
            <a href="#">Facebook</a>
            <a href="#">WhatsApp</a>
          </div>
        </div>
        <div>
          <h4 class="footer__h">Categorías</h4>
          <div class="footer__links">
            <router-link to="/catalogo?category=lenceria">Lencería</router-link>
            <router-link to="/catalogo?category=ropa-casual">Ropa Casual</router-link>
            <router-link to="/catalogo?category=pijamas-loungewear">Pijamas & Loungewear</router-link>
            <router-link to="/catalogo?category=accesorios">Accesorios</router-link>
          </div>
        </div>
        <div>
          <h4 class="footer__h">Información</h4>
          <div class="footer__links">
            <a href="#">Guía de Talles</a>
            <a href="#">Formas de Pago</a>
            <a href="#">Envíos y Devoluciones</a>
            <a href="#">Preguntas Frecuentes</a>
          </div>
        </div>
        <div>
          <h4 class="footer__h">Pagos</h4>
          <p class="footer__pay">💳 Tarjetas de crédito y débito<br/>🏦 Transferencia bancaria<br/>📱 MercadoPago</p>
          <div class="footer__ship">🚚 Envío gratis en compras mayores a $30.000</div>
        </div>
      </div>
      <div class="footer__bottom">
        <p>© {{ new Date().getFullYear() }} Bella Boutique. Todos los derechos reservados.</p>
        <p>Hecho con ♡ en Argentina</p>
      </div>
    </div>
  </footer>
</template>

<style scoped>
.footer{background:var(--c-bg-2);border-top:1px solid var(--c-border);padding:var(--sp-16) 0 var(--sp-8);margin-top:var(--sp-24);}
.footer__grid{display:grid;grid-template-columns:2fr 1fr 1fr 1fr;gap:var(--sp-8);margin-bottom:var(--sp-12);}
@media(max-width:768px){.footer__grid{grid-template-columns:1fr 1fr;}}
@media(max-width:480px){.footer__grid{grid-template-columns:1fr;}}
.footer__brand{font-family:var(--font-display);font-size:1.5rem;font-weight:600;color:var(--c-rose);}
.footer__tag{font-size:0.85rem;color:var(--c-text-muted);margin-top:var(--sp-3);line-height:1.6;}
.footer__social{display:flex;flex-direction:column;gap:var(--sp-2);margin-top:var(--sp-4);}
.footer__social a{font-size:0.85rem;color:var(--c-text-muted);transition:color var(--t-fast);}
.footer__social a:hover{color:var(--c-rose);}
.footer__h{font-family:var(--font-display);font-size:1rem;font-weight:500;margin-bottom:var(--sp-4);}
.footer__links{display:flex;flex-direction:column;gap:var(--sp-3);}
.footer__links a{font-size:0.85rem;color:var(--c-text-muted);transition:color var(--t-fast);}
.footer__links a:hover{color:var(--c-rose);}
.footer__pay{font-size:0.85rem;color:var(--c-text-muted);line-height:2;}
.footer__ship{margin-top:var(--sp-4);padding:var(--sp-3);background:rgba(201,132,122,0.1);border-radius:var(--r-md);border:1px solid var(--c-border-rose);font-size:0.8rem;color:var(--c-text-muted);}
.footer__bottom{display:flex;justify-content:space-between;align-items:center;padding-top:var(--sp-8);border-top:1px solid var(--c-border);font-size:0.8rem;color:var(--c-text-subtle);}
@media(max-width:480px){.footer__bottom{flex-direction:column;gap:var(--sp-2);text-align:center;}}
</style>
