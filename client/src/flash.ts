// Content managed by Project Forge, see [projectforge.md] for details.
export function flashInit() {
  const container = document.getElementById("flash-container")
  if (container === null) {
    return;
  }
  const x = container.querySelectorAll(".flash");
  if (x.length > 0) {
    setTimeout(() => {
      for (const f of x) {
        const el = f as HTMLElement;
        el.style.opacity = "0";
        setTimeout(() => el.remove(), 500);
      }
    }, 3000)
  }
}
