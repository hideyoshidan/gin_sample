package domain

type Samples struct {
    ID int
	Text string
    CreatedAt int64
    UpdatedAt int64
}

// この struct はビジネスロジックだと思うので、 usecase で書くべきなのかと思ったけど、
// ここに定義した。
type SamplesForGet struct {
    ID int `json:"id"`
    Text string `json:"text"`
}

func (u *Samples) BuildForGet() SamplesForGet {
    sample := SamplesForGet{}
    sample.ID = u.ID
    sample.Text = u.Text

    return sample
}