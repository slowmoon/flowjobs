package opts

import "time"

type Option func(options *Options)error


type Options struct {
	Limit int
	TTL time.Duration
	Other map[interface{}]interface{}
}

func (ops *Options)Apply(opt...Option)error  {
    for _, o := range opt {
    	if err :=  o(ops);err!= nil {
    	   return  err
		}
	}
    return  nil
}

func LimitOption(limit int) Option {
  return func(options *Options) error {
  	options.Limit = limit
  	return  nil
  }
}

func TTLOption(ttl time.Duration)Option  {
   return func(options *Options) error {
	  options.TTL = ttl
	  return  nil
   }
}
