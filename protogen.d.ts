export class Template {
    constructor(file: string);

    TemplateName(): string;

    Execute(): string;

    Set(name: string, val: any);

    RegisterFunc(name: string, fn: (...params: any) => {});
}

export interface Context {
    Packages(): Package[];

    Package(name: string): Package;

    GlobalScope(): Package;
}

export interface Generator {
    Process(ctx: Context): void;

    GetFiles(): File[];
}

export interface Options {
    Has(name: string): boolean;

    Get(name: string): any;
}

export interface Package extends Options {
    Name(): string;

    Services(): Service[];

    Service(name: string): Service;

    Messages(): Message[];

    Message(name: string): Message;

    Enums(): Enum[];

    Enum(name: string): Enum;

}

export interface Service extends Options {
    Name(): string;

    RPCs(): RPC[];

    RPC(name: string): RPC;
}

export interface RPC extends Options {
    Name(): string;

    RequestMessageType(): Type;

    ResponseMessageType(): Type;

    // IsRequestStreaming: boolean;
    // IsResponseStreaming: boolean;
}

export interface Message extends Options {
    Name(): string;

    Type(): Type;

    Enums(): Enum[];

    Enum(name: string): Enum;

    Messages(): Message[];

    Message(name: string): Message;

    IsOneOf(): boolean;

    Fields(): Field[];

    Field(name: string): Field;
}

export interface Field extends Options {
    Name(): string;

    Type(): Type;

    FieldNumber(): string;

    IsRepeated(): boolean;

}

export interface Type {
    IsNative(): boolean;

    GetNativeType(): string;

    GetFullyQualifiedName(): string;

    IsMessage(): boolean;

    GetMessage(): Message;

    IsEnum(): boolean;

    GetEnum(): Enum;
}

export type File = {
    Name: string;
    Contents: string;
}

export interface Enum {
    Name(): string;

    Values(): EnumValue;

}

export interface EnumValue {
    Identifier(): string;

    Number(): string;
}