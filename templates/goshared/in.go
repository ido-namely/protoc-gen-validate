package goshared

const inTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.In }}
		if _, ok := {{ lookup $f "InLookup" }}[{{ accessor . }}]; !ok {
			err := {{ err . "value must be in list " $r.In }}
			if stopOnError { return err }
			multiErr = multierror.Append(multiErr, err)
		}
	{{ else if $r.NotIn }}
		if _, ok := {{ lookup $f "NotInLookup" }}[{{ accessor . }}]; ok {
			err := {{ err . "value must not be in list " $r.NotIn }}
			if stopOnError { return err }
			multiErr = multierror.Append(multiErr, err)
		}
	{{ end }}
`
