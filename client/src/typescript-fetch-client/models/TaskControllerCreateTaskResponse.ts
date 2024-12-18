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
import type { ModelsTask } from './ModelsTask';
import {
    ModelsTaskFromJSON,
    ModelsTaskFromJSONTyped,
    ModelsTaskToJSON,
    ModelsTaskToJSONTyped,
} from './ModelsTask';

/**
 * 
 * @export
 * @interface TaskControllerCreateTaskResponse
 */
export interface TaskControllerCreateTaskResponse {
    /**
     * 
     * @type {ModelsTask}
     * @memberof TaskControllerCreateTaskResponse
     */
    task?: ModelsTask;
}

/**
 * Check if a given object implements the TaskControllerCreateTaskResponse interface.
 */
export function instanceOfTaskControllerCreateTaskResponse(value: object): value is TaskControllerCreateTaskResponse {
    return true;
}

export function TaskControllerCreateTaskResponseFromJSON(json: any): TaskControllerCreateTaskResponse {
    return TaskControllerCreateTaskResponseFromJSONTyped(json, false);
}

export function TaskControllerCreateTaskResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): TaskControllerCreateTaskResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'task': json['task'] == null ? undefined : ModelsTaskFromJSON(json['task']),
    };
}

export function TaskControllerCreateTaskResponseToJSON(json: any): TaskControllerCreateTaskResponse {
    return TaskControllerCreateTaskResponseToJSONTyped(json, false);
}

export function TaskControllerCreateTaskResponseToJSONTyped(value?: TaskControllerCreateTaskResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'task': ModelsTaskToJSON(value['task']),
    };
}

