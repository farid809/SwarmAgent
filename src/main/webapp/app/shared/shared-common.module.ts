import { NgModule } from '@angular/core';

import { SwarmAgentSharedLibsModule, JhiAlertComponent, JhiAlertErrorComponent } from './';

@NgModule({
    imports: [SwarmAgentSharedLibsModule],
    declarations: [JhiAlertComponent, JhiAlertErrorComponent],
    exports: [SwarmAgentSharedLibsModule, JhiAlertComponent, JhiAlertErrorComponent]
})
export class SwarmAgentSharedCommonModule {}
