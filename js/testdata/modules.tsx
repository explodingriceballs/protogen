import {ModuleClass} from "./modules_export";

export default class TestClass {
    public TestFn(): string {
        let m = new ModuleClass("foobar");
        return m.GetString();
    }
}