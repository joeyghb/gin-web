package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type StructA struct {
    FieldA string `form:"field_a"`
}

type StructB struct {
    NestedStruct StructA
    FieldB string `form:"field_b"`
}

type StructC struct {
    NestedStructPointer *StructA
    FieldC string `form:"field_c"`
}

type StructD struct {
    NestedAnonyStruct struct {
        FieldX string `form:"field_x"`
    }
    FieldD string `form:"field_d"`
}

//$ curl "http://localhost:8080/getb?field_a=hello&field_b=world"
//{"a":{"FieldA":"hello"},"b":"world"}
func GetDataB(c *gin.Context) {
    var b StructB
    c.Bind(&b)
    c.JSON(http.StatusOK, gin.H{
        "a": b.NestedStruct,
        "b": b.FieldB,
    })
}

//$ curl "http://localhost:8080/getc?field_a=hello&field_c=world"
//{"a":{"FieldA":"hello"},"c":"world"}
func GetDataC(c *gin.Context) {
    var b StructC
    c.Bind(&b)
    c.JSON(http.StatusOK, gin.H{
        "a": b.NestedStructPointer,
        "c": b.FieldC,
    })
}

//$ curl "http://localhost:8080/getd?field_x=hello&field_d=world"
//{"d":"world","x":{"FieldX":"hello"}}
func GetDataD(c *gin.Context) {
    var b StructD
    c.Bind(&b)
    c.JSON(http.StatusOK, gin.H{
        "x": b.NestedAnonyStruct,
        "d": b.FieldD,
    })
}