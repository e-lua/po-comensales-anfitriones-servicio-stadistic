package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
)

func Pg_Insert_OrderDetails(orderdetails []models.Pg_Element) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var insumos_pg []interface{}

	//Instanciando los valores
	idelement_pg, idbusiness_pg, idorder_pg, name_pg, idcarta_pg, url_pg, description_pg, typemoney_pg, unitprice_pg, quantity_pg, discount_pg, category_pg, typefood_pg, idcategory_pg, costos_pg := []int{}, []int{}, []int64{}, []string{}, []int{}, []string{}, []string{}, []int{}, []float64{}, []int{}, []float32{}, []string{}, []string{}, []int{}, []float64{}

	for _, od := range orderdetails {
		idelement_pg = append(idelement_pg, od.IDElement)
		idbusiness_pg = append(idbusiness_pg, od.IDBusiness)
		idorder_pg = append(idorder_pg, od.IDOrder)
		name_pg = append(name_pg, od.NameE)
		idcarta_pg = append(idcarta_pg, od.IDCarta)
		url_pg = append(url_pg, od.URLPhoto)
		description_pg = append(description_pg, od.Description)
		typemoney_pg = append(typemoney_pg, od.TypeMoney)
		unitprice_pg = append(unitprice_pg, od.UnitPrice)
		quantity_pg = append(quantity_pg, od.Quantity)
		discount_pg = append(discount_pg, od.Discount)
		category_pg = append(category_pg, od.Category)
		typefood_pg = append(typefood_pg, od.Typefood)
		idcategory_pg = append(idcategory_pg, od.IdCategory)
		insumos_pg = append(insumos_pg, od.Insumos)
		costos_pg = append(costos_pg, od.Costos)
	}

	//Enviado los datos a la base de datos
	db := models.Conectar_Pg_DB()

	query := `INSERT INTO OrderDetails(idelement,idorder,idbusiness,idcarta,unitprice,quantity,discount,namee,descriptione,typemoney,urle,category,typefood,idcategory,insumos,costo) (select * from unnest($1::int[], $2::bigint[],$3::int[],$4::int[],$5::decimal(8,2)[],$6::int[],$7::decimal(8,2)[],$8::varchar(100)[],$9::varchar(250)[],$10::int[],$11::varchar(230)[],$12::varchar(100)[],$13::varchar(100)[],$14::int[],$15::jsonb[],$16::real[]))`
	if _, err := db.Exec(ctx, query, idelement_pg, idorder_pg, idbusiness_pg, idcarta_pg, unitprice_pg, quantity_pg, discount_pg, name_pg, description_pg, typemoney_pg, url_pg, category_pg, typefood_pg, idcategory_pg, insumos_pg, costos_pg); err != nil {
		return err
	}

	return nil
}
