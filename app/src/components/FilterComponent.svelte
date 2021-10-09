<script lang="ts">
   import { Col, FormGroup, Input, Label, Row } from "sveltestrap";
</script>

<Col style="align-text:center; margin-top:1em;">
   <Row>
      
      <Col xs="2"><Input id="cAlpha" type="checkbox" label="Alpha" checked /></Col>
      <Col xs="2"><Input id="cBeta" type="checkbox" label="Beta" checked /></Col>
      <Col xs="2"><Input id="cDeprec" type="checkbox" label="Deprecated" checked /></Col>
      <Col>
         <Row>
            <Col><Label style="text-align: right;">cut off matches under</Label></Col>
            <Col><Input type="number" id="exampleNumber" min={0} max={100} placeholder="0%"/></Col>
         </Row>
      </Col>
   </Row>
</Col>
