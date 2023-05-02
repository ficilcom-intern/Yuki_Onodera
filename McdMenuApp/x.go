
    "encoding/csv"
    "fmt"
    "os"

    "github.com/PuerkitoBio/goquery"
)

type Product struct {
    Name            string  `gorm:"column:name"`
    Weight          string  `gorm:"column:weight"`
    Calories        float64 `gorm:"column:calories"`
    Protein         float64 `gorm:"column:protein"`
    Fat             float64 `gorm:"column:fat"`
    Carbohydrates   float64 `gorm:"column:carbohydrates"`
    Minerals        string  `gorm:"column:minerals"`
    Vitamins        string  `gorm:"column:vitamins"`
    Cholesterol     float64 `gorm:"column:cholesterol"`
    DietaryFiber    float64 `gorm:"column:dietary_fiber"`
    SaltEquivalent  float64 `gorm:"column:salt_equivalent"`
    Sodium          float64 `gorm:"column:sodium"`
    Potassium       float64 `gorm:"column:potassium"`
    Calcium         float64 `gorm:"column:calcium"`
    Phosphorus      float64 `gorm:"column:phosphorus"`
    Iron            float64 `gorm:"column:iron"`
    VitaminA        float64 `gorm:"column:vitamin_a"`
    VitaminB1       float64 `gorm:"column:vitamin_b1"`
    VitaminB2       float64 `gorm:"column:vitamin_b2"`
    Niacin          float64 `gorm:"column:niacin"`
    VitaminC        float64 `gorm:"column:vitamin_c"`
}


func main() {
    url := "https://www.mcdonalds.co.jp/quality/allergy_Nutrition/nutrient/"

    doc, err := goquery.NewDocument(url)
    if err != nil {
        panic(err)
    }

    var products []Product

    doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
        if i == 0 {
            // ヘッダー行はスキップする
            return
        }

        productName := s.Find("th").Text()
        productWeight := s.Find("td").Eq(12).Text()
        totalCalories, _ := strconv.ParseFloat(s.Find("td").Eq(0).Text(), 64)
        totalProtein, _ := strconv.ParseFloat(s.Find("td").Eq(1).Text(), 64)
        totalFat, _ := strconv.ParseFloat(s.Find("td").Eq(2).Text(), 64)
        saturatedFat, _ := strconv.ParseFloat(s.Find("td").Eq(3).Text(), 64)
        transFat, _ := strconv.ParseFloat(s.Find("td").Eq(4).Text(), 64)
        cholesterol, _ := strconv.ParseFloat(s.Find("td").Eq(5).Text(), 64)
        sodium, _ := strconv.ParseFloat(s.Find("td").Eq(6).Text(), 64)
        totalCarbohydrates, _ := strconv.ParseFloat(s.Find("td").Eq(7).Text(), 64)
        dietaryFiber, _ := strconv.ParseFloat(s.Find("td").Eq(8).Text(), 64)
        sugars, _ := strconv.ParseFloat(s.Find("td").Eq(9).Text(), 64)
        vitaminA, _ := strconv.ParseFloat(s.Find("td").Eq(10).Text(), 64)
        vitaminB1, _ := strconv.ParseFloat(s.Find("td").Eq(19).Text(), 64)
        vitaminB2, _ := strconv.ParseFloat(s.Find("td").Eq(20).Text(), 64)
        niacin, _ := strconv.ParseFloat(s.Find("td").Eq(21).Text(), 64)
        vitaminC, _ := strconv.ParseFloat(s.Find("td").Eq(22).Text(), 64)
        calcium, _ := strconv.ParseFloat(s.Find("td").Eq(23).Text(), 64)
        iron, _ := strconv.ParseFloat(s.Find("td").Eq(24).Text(), 64)

        product := Product{
            Name:            productName,
            Weight:          productWeight,
            Calories:        totalCalories,
            Protein:         totalProtein,
            Fat:             totalFat,
            SaturatedFat
