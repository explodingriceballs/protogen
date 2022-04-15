// The definitions are in the protogen.d.ts file
import {File, Generator, TemplateEngine, Package} from "../../protogen";

// Export the class
export default class MyGenerator implements Generator {
    private templateEngine: TemplateEngine;
    private generatedFiles: File[] = [];

    // The template engine is passed when the generator class is
    // instantiated
    public constructor(t: TemplateEngine) {
        this.templateEngine = t;
    }

    public getFiles(): File[] {
        return this.generatedFiles;
    }

    public processPackage(p: Package): void {
        const renderedFile = this.templateEngine.render("sample.html.tmpl", {
            Package: p,
        });
        this.generatedFiles.push({
            Name: "index.html",
            Contents: renderedFile,
        })
    }
}