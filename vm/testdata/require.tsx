import {TestEnum} from "./required_enum";
import TestRequiredClass from "./required_class";

export default class TestClass {
    public TestMethod(): number {
        return TestEnum.Hello;
    }

    public TestMethod2(): number {
        let testClass = new TestRequiredClass();
        return testClass.TestFunc();
    }
}