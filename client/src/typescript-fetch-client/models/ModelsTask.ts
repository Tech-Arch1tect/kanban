/* tslint:disable */
/* eslint-disable */
/**
 * Server API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { GormDeletedAt } from './GormDeletedAt';
import {
    GormDeletedAtFromJSON,
    GormDeletedAtFromJSONTyped,
    GormDeletedAtToJSON,
    GormDeletedAtToJSONTyped,
} from './GormDeletedAt';
import type { ModelsComment } from './ModelsComment';
import {
    ModelsCommentFromJSON,
    ModelsCommentFromJSONTyped,
    ModelsCommentToJSON,
    ModelsCommentToJSONTyped,
} from './ModelsComment';
import type { ModelsSwimlane } from './ModelsSwimlane';
import {
    ModelsSwimlaneFromJSON,
    ModelsSwimlaneFromJSONTyped,
    ModelsSwimlaneToJSON,
    ModelsSwimlaneToJSONTyped,
} from './ModelsSwimlane';

/**
 * 
 * @export
 * @interface ModelsTask
 */
export interface ModelsTask {
    /**
     * 
     * @type {number}
     * @memberof ModelsTask
     */
    boardId?: number;
    /**
     * 
     * @type {Array<ModelsComment>}
     * @memberof ModelsTask
     */
    comments?: Array<ModelsComment>;
    /**
     * 
     * @type {string}
     * @memberof ModelsTask
     */
    createdAt?: string;
    /**
     * 
     * @type {GormDeletedAt}
     * @memberof ModelsTask
     */
    deletedAt?: GormDeletedAt;
    /**
     * 
     * @type {string}
     * @memberof ModelsTask
     */
    description?: string;
    /**
     * 
     * @type {number}
     * @memberof ModelsTask
     */
    id?: number;
    /**
     * 
     * @type {string}
     * @memberof ModelsTask
     */
    status?: string;
    /**
     * 
     * @type {ModelsSwimlane}
     * @memberof ModelsTask
     */
    swimlane?: ModelsSwimlane;
    /**
     * 
     * @type {number}
     * @memberof ModelsTask
     */
    swimlaneId?: number;
    /**
     * 
     * @type {string}
     * @memberof ModelsTask
     */
    title?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelsTask
     */
    updatedAt?: string;
}

/**
 * Check if a given object implements the ModelsTask interface.
 */
export function instanceOfModelsTask(value: object): value is ModelsTask {
    return true;
}

export function ModelsTaskFromJSON(json: any): ModelsTask {
    return ModelsTaskFromJSONTyped(json, false);
}

export function ModelsTaskFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelsTask {
    if (json == null) {
        return json;
    }
    return {
        
        'boardId': json['board_id'] == null ? undefined : json['board_id'],
        'comments': json['comments'] == null ? undefined : ((json['comments'] as Array<any>).map(ModelsCommentFromJSON)),
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'deletedAt': json['deleted_at'] == null ? undefined : GormDeletedAtFromJSON(json['deleted_at']),
        'description': json['description'] == null ? undefined : json['description'],
        'id': json['id'] == null ? undefined : json['id'],
        'status': json['status'] == null ? undefined : json['status'],
        'swimlane': json['swimlane'] == null ? undefined : ModelsSwimlaneFromJSON(json['swimlane']),
        'swimlaneId': json['swimlane_id'] == null ? undefined : json['swimlane_id'],
        'title': json['title'] == null ? undefined : json['title'],
        'updatedAt': json['updated_at'] == null ? undefined : json['updated_at'],
    };
}

export function ModelsTaskToJSON(json: any): ModelsTask {
    return ModelsTaskToJSONTyped(json, false);
}

export function ModelsTaskToJSONTyped(value?: ModelsTask | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'board_id': value['boardId'],
        'comments': value['comments'] == null ? undefined : ((value['comments'] as Array<any>).map(ModelsCommentToJSON)),
        'created_at': value['createdAt'],
        'deleted_at': GormDeletedAtToJSON(value['deletedAt']),
        'description': value['description'],
        'id': value['id'],
        'status': value['status'],
        'swimlane': ModelsSwimlaneToJSON(value['swimlane']),
        'swimlane_id': value['swimlaneId'],
        'title': value['title'],
        'updated_at': value['updatedAt'],
    };
}
