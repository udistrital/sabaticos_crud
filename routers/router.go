// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/sabaticos_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/estado_solicitud",
			beego.NSInclude(
				&controllers.EstadoSolicitudController{},
			),
		),

		beego.NSNamespace("/historial_solicitud",
			beego.NSInclude(
				&controllers.HistorialSolicitudController{},
			),
		),

		beego.NSNamespace("/solicitud",
			beego.NSInclude(
				&controllers.SolicitudController{},
			),
		),

		beego.NSNamespace("/soporte_solicitud",
			beego.NSInclude(
				&controllers.SoporteSolicitudController{},
			),
		),

		beego.NSNamespace("/tipo_solicitud",
			beego.NSInclude(
				&controllers.TipoSolicitudController{},
			),
		),

		beego.NSNamespace("/sabatico",
			beego.NSInclude(
				&controllers.SabaticoController{},
			),
		),

		beego.NSNamespace("/soporte_sabatico",
			beego.NSInclude(
				&controllers.SoporteSabaticoController{},
			),
		),

		beego.NSNamespace("/estado_sabatico",
			beego.NSInclude(
				&controllers.EstadoSabaticoController{},
			),
		),

		beego.NSNamespace("/estado_soporte_solicitud",
			beego.NSInclude(
				&controllers.EstadoSoporteSolicitudController{},
			),
		),

		beego.NSNamespace("/formulario_solicitud",
			beego.NSInclude(
				&controllers.FormularioSolicitudController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
