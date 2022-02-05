// Content managed by Project Forge, see [projectforge.md] for details.
import "./client.css"
import {menuInit} from "./menu";
import {flashInit} from "./flash";
import {editorInit} from "./editor";
import {linkInit} from "./link";
import {themeInit} from "./theme";
import {modeInit} from "./mode";
import {appInit} from "./app";

export function init(): void {
  (window as any).pftest = {};
  menuInit();
  modeInit();
  flashInit();
  linkInit();
  editorInit();
  themeInit();
  appInit();
}

document.addEventListener("DOMContentLoaded", init);
