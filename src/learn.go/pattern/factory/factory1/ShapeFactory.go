package factory1

type ShapeFactory struct {
}

func (this ShapeFactory) GetShape(shapeType string) Shape {
	if shapeType == "" {
		return nil
	}

	if shapeType == "1" {
		return Circle{}
	} else if shapeType == "2" {
		return Rectangle{}
	} else if shapeType == "3" {
		return Square{}
	}
	return nil
}
