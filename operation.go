package mgm

import (
	"github.com/Kamva/mgm/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func create(c *Collection, model Model) error {
	// Call to saving hook
	if err := callToBeforeCreateHooks(model); err != nil {
		return err
	}

	res, err := c.InsertOne(ctx(), model)

	if err != nil {
		return err
	}

	// Set new id
	model.SetID(res.InsertedID)

	return callToAfterCreateHooks(model)
}

func first(c *Collection, filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	return c.FindOne(ctx(), filter, opts...).Decode(model)
}

func update(c *Collection, model Model) error {
	// Call to saving hook
	if err := callToBeforeUpdateHooks(model); err != nil {
		return err
	}

	res, err := c.UpdateOne(ctx(), bson.M{field.ID: model.GetID()}, bson.M{"$set": model})

	if err != nil {
		return err
	}

	return callToAfterUpdateHooks(res, model)
}

func del(c *Collection, model Model) error {
	if err := callToBeforeDeleteHooks(model); err != nil {
		return err
	}
	res, err := c.DeleteOne(ctx(), bson.M{field.ID: model.GetID()})
	if err != nil {
		return err
	}

	return callToAfterDeleteHooks(res, model)
}
