/*
a class to statically store properties of a gcp role
mainly to be used by a RoleCollection
*/
export class Role {
    // these come from backend
    public readonly description: string;
    public readonly name: string;
    public readonly stage: string;
    public readonly title: string;
    public readonly includedPermissions: Array<string>;
    public readonly id: number;
    public matches: number
    public matchedBy: string[]
    public perc_match: number

    constructor(item: any) {
        // this comes from server side
        Object.assign(this, item)
   }
}