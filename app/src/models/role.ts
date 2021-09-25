/*
a class to statically store properties of a gcp role
mainly to be used by a RoleCollection
*/
export class Role {

        // example:
        // "description": "Ability to view or act on access approval requests and view configuration",
        // "includedPermissions": [
        //   "accessapproval.requests.approve",
        //   "accessapproval.requests.dismiss",
        //   "accessapproval.requests.get",
        //   "accessapproval.requests.list",
        //   "accessapproval.settings.get",
        //   "resourcemanager.projects.get",
        //   "resourcemanager.projects.list"
        // ],
        // "name": "roles/accessapproval.approver",
        // "stage": "BETA",
        // "title": "Access Approval Approver"

    // these come from backend
    public readonly description: string;
    public readonly name: string;
    public readonly stage: string;
    public readonly title: string;
    public readonly includedPermissions: Array<string>;
    public readonly id: number;

    // these are computed
    public matches: number
    public matchedBy: string[]
    public perc_match: number

    constructor(item: any) {
        // this comes from server side
        Object.assign(this, item)

        //client side compiled
        this.matches = 0
        this.matchedBy = []
        this.perc_match = 0
    }

    resetMatches() {
        this.matchedBy = []
        this.matches = 0
        this.perc_match = 0
    }
}