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

	var insumos_od []interface{}

	//Instanciando los valores
	idelement_od, idbusiness_od, idorder_od, name_od, idcarta_od, url_od, description_od, typemoney_od, unitprice_od, quantity_od, discount_od, category_od, typefood_od, idcategory_od, costos_od := []int{}, []int{}, []int64{}, []string{}, []int{}, []string{}, []string{}, []int{}, []float64{}, []int{}, []float32{}, []string{}, []string{}, []int{}, []float64{}

	for _, od := range orderdetails {
		idelement_od = append(idelement_od, od.IDElement)
		idbusiness_od = append(idbusiness_od, od.IDBusiness)
		idorder_od = append(idorder_od, od.IDOrder)
		name_od = append(name_od, od.NameE)
		idcarta_od = append(idcarta_od, od.IDCarta)
		url_od = append(url_od, od.URLPhoto)
		description_od = append(description_od, od.Description)
		typemoney_od = append(typemoney_od, od.TypeMoney)
		unitprice_od = append(unitprice_od, od.UnitPrice)
		quantity_od = append(quantity_od, od.Quantity)
		discount_od = append(discount_od, od.Discount)
		category_od = append(category_od, od.Category)
		typefood_od = append(typefood_od, od.Typefood)
		idcategory_od = append(idcategory_od, od.IdCategory)
		insumos_od = append(insumos_od, od.Insumos)
		costos_od = append(costos_od, od.Costos)
	}

	//Enviado los datos a la base de datos
	db := models.Conectar_Pg_DB()

	query := `INSERT INTO OrderDetails(idelement,idorder,idbusiness,idcarta,unitprice,quantity,discount,namee,descriptione,typemoney,urle,category,typefood,idcategory,insumos,costo) (select * from unnest($1::int[], $2::bigint[],$3::int[],$4::int[],$5::decimal(8,2)[],$6::int[],$7::decimal(8,2)[],$8::varchar(100)[],$9::varchar(250)[],$10::int[],$11::varchar(230)[],$12::varchar(100)[],$13::varchar(100)[],$14::int[],$15::jsonb[],$16::real[]))`
	if _, err := db.Exec(ctx, query, idelement_od, idorder_od, idbusiness_od, idcarta_od, unitprice_od, quantity_od, discount_od, name_od, description_od, typemoney_od, url_od, category_od, typefood_od, idcategory_od, insumos_od, costos_od); err != nil {
		return err
	}

	return nil
}
