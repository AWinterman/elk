{{ define "header" }}
    // Code generated by entc, DO NOT EDIT.

    {{ $pkg := base $.Config.Package }}
    {{ if hasField $ "Scope" }}
        {{ $pkg = $.Scope.Package }}
    {{ end }}
    package {{ $pkg }}
{{ end }}