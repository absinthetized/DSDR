<script lang="ts">
   import { Col, Input, Label, Row } from "sveltestrap";

   function updatePerc(e: Event) {
      let precheckValue = parseInt((e.target as HTMLInputElement).value, 10);
      const inputbox = document.getElementById("cPerc")

      if (precheckValue < 0 || precheckValue > 100) {
         inputbox.classList.add("is-invalid")

      } else {
         inputbox.classList.remove("is-invalid")
         doMinPerc = precheckValue/100
      }
   }

   export let doAlpha: Boolean = true
   export let doBeta: Boolean = true
   export let doDeprec: Boolean = true
   export let doMinPerc: Number = 0
</script>

<Col style="align-text:center; margin-top:1em;">
   <Row>
      
      <Col xs="2"><Input id="cAlpha" type="checkbox" label="Alpha" checked on:change={() => doAlpha = !doAlpha}/></Col>
      <Col xs="2"><Input id="cBeta" type="checkbox" label="Beta" checked on:change={() => doBeta = !doBeta}/></Col>
      <Col xs="2"><Input id="cDeprec" type="checkbox" label="Deprecated" checked on:change={() => doDeprec = !doDeprec}/></Col>
      <Col>
         <Row>
            <Col><Label style="text-align: right;">cut off matches under</Label></Col>
            <Col><Input type="number" id="cPerc" min={0} max={100} 
                  placeholder="0%" feedback="range 0-100%" on:change={updatePerc}/></Col>
         </Row>
      </Col>
   </Row>
</Col>
