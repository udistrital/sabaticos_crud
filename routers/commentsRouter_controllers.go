package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSabaticoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSabaticoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSabaticoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSabaticoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSabaticoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSabaticoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSabaticoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSabaticoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSabaticoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSabaticoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSolicitudController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSolicitudController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSolicitudController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSolicitudController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:EstadoSoporteSolicitudController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:FormularioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:FormularioSolicitudController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:FormularioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:FormularioSolicitudController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:FormularioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:FormularioSolicitudController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:FormularioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:FormularioSolicitudController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:FormularioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:FormularioSolicitudController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSabaticoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSabaticoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSabaticoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSabaticoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSabaticoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSolicitudController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSolicitudController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSolicitudController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSolicitudController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:HistorialSolicitudController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SabaticoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SabaticoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SabaticoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SabaticoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SabaticoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SolicitudController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SolicitudController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SolicitudController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SolicitudController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SolicitudController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSabaticoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSabaticoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSabaticoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSabaticoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSabaticoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSolicitudController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSolicitudController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSolicitudController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSolicitudController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:SoporteSolicitudController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_crud/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
