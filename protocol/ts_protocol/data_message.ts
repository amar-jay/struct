/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.15.8
 * source: data_message.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as pb_1 from "google-protobuf";
export namespace go_protocol {
    export enum DataMessageType {
        BOT = 0,
        USER = 1
    }
    export class DataMessageBody extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {}) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") { }
        }
        static fromObject(data: {}): DataMessageBody {
            const message = new DataMessageBody({});
            return message;
        }
        toObject() {
            const data: {} = {};
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): DataMessageBody {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new DataMessageBody();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): DataMessageBody {
            return DataMessageBody.deserialize(bytes);
        }
    }
    export class DataMessage extends pb_1.Message {
        #one_of_decls: number[][] = [[2]];
        constructor(data?: any[] | ({
            type?: DataMessageType;
            room_id?: string;
            room_sid?: string;
            to?: string;
            body?: DataMessageBody;
        } & (({
            message_id?: string;
        })))) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("type" in data && data.type != undefined) {
                    this.type = data.type;
                }
                if ("message_id" in data && data.message_id != undefined) {
                    this.message_id = data.message_id;
                }
                if ("room_id" in data && data.room_id != undefined) {
                    this.room_id = data.room_id;
                }
                if ("room_sid" in data && data.room_sid != undefined) {
                    this.room_sid = data.room_sid;
                }
                if ("to" in data && data.to != undefined) {
                    this.to = data.to;
                }
                if ("body" in data && data.body != undefined) {
                    this.body = data.body;
                }
            }
        }
        get type() {
            return pb_1.Message.getFieldWithDefault(this, 1, DataMessageType.BOT) as DataMessageType;
        }
        set type(value: DataMessageType) {
            pb_1.Message.setField(this, 1, value);
        }
        get message_id() {
            return pb_1.Message.getFieldWithDefault(this, 2, "") as string;
        }
        set message_id(value: string) {
            pb_1.Message.setOneofField(this, 2, this.#one_of_decls[0], value);
        }
        get has_message_id() {
            return pb_1.Message.getField(this, 2) != null;
        }
        get room_id() {
            return pb_1.Message.getFieldWithDefault(this, 3, "") as string;
        }
        set room_id(value: string) {
            pb_1.Message.setField(this, 3, value);
        }
        get room_sid() {
            return pb_1.Message.getFieldWithDefault(this, 4, "") as string;
        }
        set room_sid(value: string) {
            pb_1.Message.setField(this, 4, value);
        }
        get to() {
            return pb_1.Message.getFieldWithDefault(this, 5, "") as string;
        }
        set to(value: string) {
            pb_1.Message.setField(this, 5, value);
        }
        get body() {
            return pb_1.Message.getWrapperField(this, DataMessageBody, 6) as DataMessageBody;
        }
        set body(value: DataMessageBody) {
            pb_1.Message.setWrapperField(this, 6, value);
        }
        get has_body() {
            return pb_1.Message.getField(this, 6) != null;
        }
        get _message_id() {
            const cases: {
                [index: number]: "none" | "message_id";
            } = {
                0: "none",
                2: "message_id"
            };
            return cases[pb_1.Message.computeOneofCase(this, [2])];
        }
        static fromObject(data: {
            type?: DataMessageType;
            message_id?: string;
            room_id?: string;
            room_sid?: string;
            to?: string;
            body?: ReturnType<typeof DataMessageBody.prototype.toObject>;
        }): DataMessage {
            const message = new DataMessage({});
            if (data.type != null) {
                message.type = data.type;
            }
            if (data.message_id != null) {
                message.message_id = data.message_id;
            }
            if (data.room_id != null) {
                message.room_id = data.room_id;
            }
            if (data.room_sid != null) {
                message.room_sid = data.room_sid;
            }
            if (data.to != null) {
                message.to = data.to;
            }
            if (data.body != null) {
                message.body = DataMessageBody.fromObject(data.body);
            }
            return message;
        }
        toObject() {
            const data: {
                type?: DataMessageType;
                message_id?: string;
                room_id?: string;
                room_sid?: string;
                to?: string;
                body?: ReturnType<typeof DataMessageBody.prototype.toObject>;
            } = {};
            if (this.type != null) {
                data.type = this.type;
            }
            if (this.message_id != null) {
                data.message_id = this.message_id;
            }
            if (this.room_id != null) {
                data.room_id = this.room_id;
            }
            if (this.room_sid != null) {
                data.room_sid = this.room_sid;
            }
            if (this.to != null) {
                data.to = this.to;
            }
            if (this.body != null) {
                data.body = this.body.toObject();
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.type != DataMessageType.BOT)
                writer.writeEnum(1, this.type);
            if (this.has_message_id)
                writer.writeString(2, this.message_id);
            if (this.room_id.length)
                writer.writeString(3, this.room_id);
            if (this.room_sid.length)
                writer.writeString(4, this.room_sid);
            if (this.to.length)
                writer.writeString(5, this.to);
            if (this.has_body)
                writer.writeMessage(6, this.body, () => this.body.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): DataMessage {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new DataMessage();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.type = reader.readEnum();
                        break;
                    case 2:
                        message.message_id = reader.readString();
                        break;
                    case 3:
                        message.room_id = reader.readString();
                        break;
                    case 4:
                        message.room_sid = reader.readString();
                        break;
                    case 5:
                        message.to = reader.readString();
                        break;
                    case 6:
                        reader.readMessage(message.body, () => message.body = DataMessageBody.deserialize(reader));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): DataMessage {
            return DataMessage.deserialize(bytes);
        }
    }
}