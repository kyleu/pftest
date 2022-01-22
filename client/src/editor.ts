// Content managed by Project Forge, see [projectforge.md] for details.
const selected = "--selected";

export function editorInit() {
  let editorCache: { [key: string]: string; } = {};
  let selectedCache: { [key: string]: HTMLInputElement; } = {};
  for (const n of Array.from(document.querySelectorAll(".editor"))) {
    const frm = n as HTMLFormElement;
    const buildCache = () => {
      editorCache = {};
      selectedCache = {};
      for (const el of frm.elements) {
        const input = el as HTMLInputElement;
        if (input.name.length > 0) {
          if (input.name.endsWith(selected)) {
            selectedCache[input.name] = input;
          } else {
            if ((input.type !== "radio") || input.checked) {
              editorCache[input.name] = input.value;
            }
            const evt = () => {
              const cv = selectedCache[input.name + selected];
              if (cv) {
                cv.checked = editorCache[input.name] !== input.value;
              }
            };
            input.onchange = evt;
            input.onkeyup = evt;
          }
        }
      }
    }
    frm.onreset = buildCache;
    buildCache();
  }
}
