<script lang="ts">
   import { Col, Input, Label, Row } from "sveltestrap";

   function updatePerc(e: Event) {
      let precheckValue = parseInt((e.target as HTMLInputElement).value, 10);
      const inputbox = document.getElementById("cPerc")

      if (precheckValue < 0 || precheckValue > 100) {
         inputbox.classList.add("is-invalid")

      } else {
         inputbox.classList.remove("is-invalid")
         doMinPerc = precheckValue
      }
   }

   export let doAlpha: boolean
   export let doBeta: boolean
   export let doDeprec: boolean
   export let doMinPerc: Number
</script>

<Col style="align-text:center; margin-top:1em;">
   <Row>
      
      <Col xs="2"><Input id="cAlpha" type="checkbox" label="Alpha" bind:checked={doAlpha}/></Col>
      <Col xs="2"><Input id="cBeta" type="checkbox" label="Beta" bind:checked={doBeta}/></Col>
      <Col xs="2"><Input id="cDeprec" type="checkbox" label="Deprecated" bind:checked={doDeprec}/></Col>
      <Col>
         <Row>
            <Col><Label style="text-align: right;">cut off matches under</Label></Col>
            <Col><Input type="number" id="cPerc" min={0} max={100} 
                  placeholder="0%" value={doMinPerc.toString()} feedback="range 0-100%" on:change={updatePerc}/></Col>
         </Row>
      </Col>
   </Row>
</Col>
