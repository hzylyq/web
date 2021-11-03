package framework

import "mime/multipart"

type IRequest interface {
	QueryInt(string, int) (int, bool)
	QueryInt64(string, int) (int, bool)
	QueryFloat64(string, int) (float64, bool)
	QueryFloat32(string, int) (float64, bool)
	QueryBool(string, bool) (bool, bool)
	QueryString(string, string) (string, bool)
	QueryStringSlice(string, ...string) ([]string, bool)
	Query(string) interface{}

	ParamInt(string, int) (int, bool)
	ParamInt64(string, int) (int, bool)
	ParamFloat64(string, int) (float64, bool)
	ParamFloat32(string, int) (float64, bool)
	ParamBool(string, bool) (bool, bool)
	ParamString(string, string) (string, bool)
	ParamStringSlice(string, ...string) ([]string, bool)
	Param(string) interface{}

	FormInt(string, int) (int, bool)
	FormInt64(string, int) (int, bool)
	FormFloat64(string, int) (float64, bool)
	FormFloat32(string, int) (float64, bool)
	FormBool(string, bool) (bool, bool)
	FormString(string, string) (string, bool)
	FormStringSlice(string, ...string) ([]string, bool)
	FormFile(string) (*multipart.FileHeader, bool)
	Form(string) interface{}
}
