import { serveDir } from "@std/http/file-server";

Deno.serve((req: Request) => {
  return serveDir(req, {
    fsRoot: "./build",
  })
});

const win = new Deno.BrowserWindow({ title: "Belaberung" });
//const base = Deno.env.get("DENO_SERVE_ADDRESS")!;
//const port = base.split(":").pop();

win.addEventListener("close", ()=>{
  win.close()
  Deno.exit(0)
})