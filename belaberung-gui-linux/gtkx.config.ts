import { defineConfig } from "@gtkx/config";

export default defineConfig({
    libraries: ["Gtk-4.0"],
    applicationId: "dev.delfi.belaberung-gui-linux",
    deploy: {
        name: "Belaberung Gui Linux",
        summary: "A GTK4 application built with GTKX",
        description: [
            "Belaberung Gui Linux is a GTK4 and Adwaita application built with GTKX, which renders native GObject "
            + "widgets from React. Replace this paragraph with a description of what your application does.",
        ],
        categories: ["Utility"],
    },
});
