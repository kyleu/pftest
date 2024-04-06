// Content managed by Project Forge, see [projectforge.md] for details.
import * as d3 from "npm:d3";

export async function load(u, t) {
  const response = await fetch(u);
  if (!response.ok) throw new Error(`fetch failed: ${response.status}`);
  if (t === "csv") {
    return d3.csvParse(await response.text());
  }
  if (t === "json" || t === "") {
    return await response.json();
  }
  return await response.text();
}

const defaultOpts = {limit: 0, offset: 0, order: "", t: "json", q: "", extra: {}};

export function urlFor(key, path, {limit, offset, order, t, q, extra} = defaultOpts) {
  let ret = `http://localhost:41000/${path}?t=${t}`;
  let prefix = key !== "" ? key + "." : "";
  if (order && order !== "") {
    ret += `&${prefix}o=${order}`;
  }
  if (limit > 0) {
    ret += `&${prefix}l=${limit}`;
  }
  if (offset > 0) {
    ret += `&${prefix}x=${offset}`;
  }
  if (q && q !== "") {
    ret += `&q=${encodeURIComponent(q)}`;
  }
  if (extra) {
    for (const [key, value] of Object.entries(extra)) {
      ret += `&${encodeURIComponent(key)}=${encodeURIComponent(value)}`;

    }
  }
  return ret;
}

export async function capitals(opts) {
  return await load(urlFor("capital", "capital", opts), opts?.t);
}

export async function auditeds(opts) {
  return await load(urlFor("audited", "audited", opts), opts?.t);
}

export async function basics(opts) {
  return await load(urlFor("basic", "basic", opts), opts?.t);
}

export async function mixedCases(opts) {
  return await load(urlFor("mixedcase", "mixedcase", opts), opts?.t);
}

export async function paths(opts) {
  return await load(urlFor("path", "g1/g2/path", opts), opts?.t);
}

export async function references(opts) {
  return await load(urlFor("reference", "reference", opts), opts?.t);
}

export async function relations(opts) {
  return await load(urlFor("relation", "relation", opts), opts?.t);
}

export async function seeds(opts) {
  return await load(urlFor("seed", "seed", opts), opts?.t);
}

export async function softdels(opts) {
  return await load(urlFor("softdel", "softdel", opts), opts?.t);
}

export async function timestamps(opts) {
  return await load(urlFor("timestamp", "timestamp", opts), opts?.t);
}

export async function troubles(opts) {
  return await load(urlFor("trouble", "troub/le", opts), opts?.t);
}
