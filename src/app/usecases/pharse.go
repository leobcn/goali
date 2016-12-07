package usecases

import "app"

func NewPharse(r app.IDatabase) *Pharse {
	return &Pharse{r}
}

type Pharse struct {
	r app.IDatabase
}

func (p *Pharse) GetOrCreate(pm *app.Pharse) error {
	// todo: improve this fucking snip
	if err := p.r.OneBy(pm, app.DBWhere{"Sum": pm.Sum}); err != nil {
		if p.r.IsNotFoundErr(err) {
			return p.r.Store(pm)
		}
		return err
	}
	return nil
}
