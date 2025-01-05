package contacts

import (
	"errors"
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

func TestService_GetIdByName(t *testing.T) {
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
		want    int
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
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.fields.repository,
			}
			got, err := s.GetIdByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIdByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetIdByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Create(t *testing.T) {
	type fields struct {
		repository Repository
	}
	type args struct {
		contact *contact.Contact
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				contact: &contact.Contact{
					Name: "Max",
				},
			},
			fields: fields{
				repository: func() Repository {
					model := &contact.Contact{
						Name: "Max",
					}

					r := new(mocks.Repository)
					r.On("Create", model).Return(nil)

					return r
				}(),
			},
			wantErr: false,
		},
		{
			name: "error on empty name",
			args: args{
				contact: &contact.Contact{
					Name: "",
				},
			},
			fields: fields{
				repository: func() Repository {
					model := &contact.Contact{
						Name: "",
					}

					r := new(mocks.Repository)
					r.On("Create", model).Return(nil)

					return r
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.fields.repository,
			}
			if err := s.Create(tt.args.contact); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_Update(t *testing.T) {
	type fields struct {
		repository Repository
	}
	type args struct {
		name              string
		contactUpdateData *contact.ContactUpdateData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				name: "Max",
				contactUpdateData: &contact.ContactUpdateData{
					Name: "New Max",
				},
			},
			fields: fields{
				repository: func() Repository {
					model := &contact.Contact{
						Name:     "Max",
						Telegram: "Old Telegram",
					}
					updateModel := &contact.ContactUpdateData{
						Name: "New Max",
					}

					r := new(mocks.Repository)
					r.On("GetByName", "Max").Return(model, nil)
					r.On("Update", model, updateModel).Return(nil)

					return r
				}(),
			},
		},
		{
			name: "error on contact not found",
			args: args{
				name: "Max",
				contactUpdateData: &contact.ContactUpdateData{
					Name: "New Max",
				},
			},
			fields: fields{
				repository: func() Repository {
					r := new(mocks.Repository)
					r.On("GetByName", "Max").Return(nil, errors.New("some error"))

					return r
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.fields.repository,
			}
			if err := s.Update(tt.args.name, tt.args.contactUpdateData); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
