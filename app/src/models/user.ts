/**
 * a firebase user 
 */
import type { User } from "firebase/auth";

class CurrentUser {
    private _currentUser: User | null
    constructor() {
        this._currentUser = null
    }

    get info() {
        return this._currentUser
    }

    set info(aUser: User | null) {
        this._currentUser = aUser
    }
}

const currentUser = new CurrentUser()
export default currentUser