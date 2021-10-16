/* stores the logic for the filter component and let it proagate all around the app */

export class FilterService {
   public doAlpha: boolean
   public doBeta: boolean
   public doDeprec: boolean
   public doMinPerc: Number

   constructor() {
      this.doAlpha = true
      this.doBeta = true
      this.doDeprec = true
      this.doMinPerc = 0
   }
}