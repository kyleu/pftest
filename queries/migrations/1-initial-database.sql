-- $PF_IGNORE$
-- {% import "github.com/kyleu/pftest/queries/ddl" %}
-- {% func Migration1InitialDatabase(debug bool) %}

-- {%- if debug -%}
-- {%= ddl.DropAll() %}
-- {%- endif -%}

-- {%= ddl.CreateAll() %}

-- {% endfunc %}
