// Content managed by Project Forge, see [projectforge.md] for details.
export const appKey = "pftest";
export const appName = "Test Project";

export function svgRef(key: string, size: number): string {
  return `<svg class="icon" style="width: ${size}px; height: ${size}px;"><use xlink:href="#svg-${key}"></use></svg>`;
}
