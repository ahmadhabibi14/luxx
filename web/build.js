import chokidar from "chokidar";
import { exec } from "child_process";
import { join } from "path";

const buildCommand = "pnpm build";
const ignorePath = new Set([
	"node_modules",
	"dist",
	".vscode",
	".idea",
	".git",
	".gitignore",
	"build.js",
	"pnpm-lock.json",
	"package.json",
	"README.md",
	"build.js",
	"vite.config.js*"
]);
const watcher = chokidar.watch(".", {
	ignored: (s) => ignorePath.has(s) || ignorePath.has(join("./", s)),
	ignoreInitial: true
});

// Listen for change events
watcher.on("change", (path) => {
	console.log(`File ${path} has changed. Running build command...`);
	exec(buildCommand);
});
