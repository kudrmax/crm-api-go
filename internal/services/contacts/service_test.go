package contacts

import (
	"reflect"
	"testing"

	"my/crm-golang/internal/models/contact"
	"my/crm-golang/internal/my_errors"
	"my/crm-golang/internal/services/contacts/mocks"
)

func TestService_GetByName(t *testing.T) {
	type fields struct {
		repository Repository
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *contact.Contact
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				name: "Max",
			},
			fields: fields{
				repository: func() Repository {
					r := new(mocks.Repository)
					r.On("GetByName", "Max").Return(
						&contact.Contact{
							Id:   1,
							Name: "Max",
						}, nil)

					return r
				}(),
			},
			want: &contact.Contact{
				Id:   1,
				Name: "Max",
			},
		},
		{
			name: "error not found",
			args: args{
				name: "Mda",
			},
			fields: fields{
				repository: func() Repository {
					r := new(mocks.Repository)
					r.On("GetByName", "Mda").Return(
						nil, my_errors.ContactNotFoundErr)

					return r
				}(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.fields.repository,
			}
			got, err := s.GetByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
