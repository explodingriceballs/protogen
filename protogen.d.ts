export class TemplateEngine {
    render(template: string, params: any, functions?: {}): string;
}

export interface Generator {
    processPackage(p: Package): void;

    getFiles(): File[];
}

export type Package = {
    Name: string;
    Services: Service[];
    Messages: Message[];
    Enums: Enum[];
    Extensions: Message[];
    Options: Option[];
}

export type Service = {
    Name: string;
    RPCs: RPC[];
}

export type RPC = {
    Name: string;
    RequestMsg: string;
    ResponseMsg: string;
    IsRequestStreaming: boolean;
    IsResponseStreaming: boolean;
}

export type Option = {
    Key: string;
    Value: string;
}

export type Message = {
    Name: string;
    Option: Option[];
    Enums: Enum[];
    Fields: Field[];
    NestedMessages: Message[];
    OneOfs: OneOf[];
}

export type OneOf = {
    Name: string;
    Fields: Field[];
}

export type Field = {
    Name: string;
    Type: string;
    FieldNumber: string;
    IsRepeated: boolean;
    Options: Option[];
}

export type File = {
    Name: string;
    Contents: string;
}

export type Enum = {
    Name: string
    Options: Option[];
    Values: EnumValue[];
}

export type EnumValue = {
    Name: string;
    Value: string;
    Options: Option[];
}