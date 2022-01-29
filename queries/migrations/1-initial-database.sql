-- $PF_IGNORE$
-- {% import "github.com/kyleu/pftest/queries/ddl" %}
-- {% func Migration1InitialDatabase(debug bool) %}

-- {%- if debug -%}
-- {%= ddl.AuditDrop() %}
-- {%= ddl.DropAll() %}
-- {%- endif -%}

-- {%= ddl.CreateAll() %}
-- {%= ddl.AuditCreate() %}

-- {% endfunc %}
