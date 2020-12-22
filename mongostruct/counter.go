package mongostruct

type counter struct {
	ID            string `bson:"_id"`
	SequenceValue int    `bson:"sequence_value"`
}
