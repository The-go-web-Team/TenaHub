package repository

// // HcRepository implements healthcenter.Repository
// type HcRepository struct {
// 	conn *gorm.DB
// }

// // NewHcRepository creates object of HcRepository
// func NewHcRepository(conn *gorm.DB) *HcRepository {
// 	return &HcRepository{conn: conn}
// }

// HealthCenter returns single healthcenter from database
//func (hcr *HealthCenterGormRepo) SingleHealthCenter(id uint) (*entity.HealthCenter, []error) {
//	hcs := entity.HealthCenter{}
//	errs := hcr.conn.Where("id = ?", id).First(&hcs).GetErrors()
//
//	if len(errs) > 0 {
//		return nil, errs
//	}
//
//	return &hcs, nil
//}
//
//// HealthCenters returns all healthcenters data from database
//func (hcr *HealthCenterGormRepo) SearchHealthCenters(value string, column string) ([]entity.Hcrating, []error) {
//	// hcs := []entity.HealthCenter{}
//	hcsRating := []entity.Hcrating{}
//	fmt.Println("value, column", value, column)
//	switch column {
//	case "name":
//		fmt.Println("name")
//		// errs := hcr.conn.Where("name ILIKE ?", "%"+value+"%").Find(&hcs).GetErrors()
//		errs := hcr.conn.Raw("select health_centers.*, avg(comments.rating) as rating from health_centers left join comments on health_centers.id = comments.health_center_id where health_centers.name ILIKE ? group by health_centers.id order by rating desc;", "%"+value+"%").Scan(&hcsRating).GetErrors()
//		fmt.Println(hcsRating)
//		if len(errs) > 0 {
//			return nil, errs
//		}
//		return hcsRating, nil
//	case "city":
//		fmt.Println("city")
//		// errs := hcr.conn.Where("city ILIKE ?", "%"+value+"%").Find(&hcs).GetErrors()
//		errs := hcr.conn.Raw("select health_centers.*, avg(comments.rating) as rating from health_centers left join comments on health_centers.id = comments.health_center_id where health_centers.city ILIKE ? group by health_centers.id order by rating desc;", "%"+value+"%").Scan(&hcsRating).GetErrors()
//
//		if len(errs) > 0 {
//			return nil, errs
//		}
//		return hcsRating, nil
//	case "service":
//		fmt.Println("service")
//		// errs := hcr.conn.Raw("select * from health_centers where id in (?)", hcr.conn.Table("services").Select("health_center_id").Where("name ILIKE ?", "%"+value+"%").QueryExpr()).Find(&hcr).GetErrors()
//		result := []struct {
//			HealthCenterID int
//		}{}
//		errs := hcr.conn.Table("services").Select("health_center_id").Where("name ILIKE ?", "%"+value+"%").Find(&result).GetErrors()
//		fmt.Println(result)
//
//		if len(errs) > 0 {
//			return nil, errs
//		}
//
//		arr := make([]int, len(result))
//
//		for _, hid := range result {
//			arr = append(arr, hid.HealthCenterID)
//		}
//		// errs = hcr.conn.Table("health_centers").Select("health_centers.*, avg(comments.rating) as rating").Joins("left join comments on comments.health_center_id = health_centers.id").Where("health_centers.id in (?)", arr).Group("healch_centers.id").Scan(&hcsRating).GetErrors()
//		errs = hcr.conn.Raw("select health_centers.*, avg(comments.rating) as rating from health_centers left join comments on health_centers.id = comments.health_center_id where health_centers.id in (?) group by health_centers.id order by rating desc;", arr).Scan(&hcsRating).GetErrors()
//
//		fmt.Println(errs)
//		fmt.Println(hcsRating)
//
//		if len(errs) > 0 {
//			return nil, errs
//		}
//		return hcsRating, nil
//	default:
//		fmt.Println("default")
//		errs := hcr.conn.Raw("select health_centers.*, avg(comments.rating) as rating from health_centers left join comments on health_centers.id = comments.health_center_id where health_centers.name ILIKE ? group by health_centers.id order by rating desc;", "%"+value+"%").Scan(&hcsRating).GetErrors()
//		if len(errs) > 0 {
//			return nil, errs
//		}
//		return hcsRating, nil
//	}
//}

// // Top returns healthcenters with rating from database
// func (hcr *HcRepository) Top(amount uint) ([]entity.Hcrating, []error) {
// 	result := []entity.Hcrating{}
// 	errs := hcr.conn.Raw("select health_centers.*, avg(comments.rating) as rating from health_centers left join comments on health_centers.id = comments.health_center_id group by health_centers.id order by rating limit ?;", amount).Scan(&result).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	fmt.Println(result)
// 	return result, nil
// }

// // SearchHealthCenter searches healthcenters by name from database
// // func (hcr *HcRepository) SearchHealthCenter(name string)([]entity.HealthCenter, []error) {
// // return nil, nil
// // }

// // UpdateHealthCenter updates healthcenter on database
// func (hcr *HcRepository) UpdateHealthCenter(hc entity.HealthCenter) (*entity.HealthCenter, []error) {
// 	return nil, nil
// }

// // StoreHealthCenter stores healthcenter data to the database
// func (hcr *HcRepository) StoreHealthCenter(hc entity.HealthCenter) (*entity.HealthCenter, []error) {
// 	return nil, nil
// }

// // DeleteHealthCenter deletes healthcenter from database
// func (hcr *HcRepository) DeleteHealthCenter(id uint) (*entity.HealthCenter, []error) {
// 	return nil, nil
// }
