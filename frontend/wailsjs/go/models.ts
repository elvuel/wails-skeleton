export namespace main {
	
	export class BootstrapData {
	    appName: string;
	    locale: string;
	    supportedLocales: string[];
	    databaseDriver: string;
	    databasePath: string;
	    databaseReady: boolean;
	    trayEnabled: boolean;
	    hideWindowOnClose: boolean;
	    message: string;
	    startupError: string;
	
	    static createFrom(source: any = {}) {
	        return new BootstrapData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.appName = source["appName"];
	        this.locale = source["locale"];
	        this.supportedLocales = source["supportedLocales"];
	        this.databaseDriver = source["databaseDriver"];
	        this.databasePath = source["databasePath"];
	        this.databaseReady = source["databaseReady"];
	        this.trayEnabled = source["trayEnabled"];
	        this.hideWindowOnClose = source["hideWindowOnClose"];
	        this.message = source["message"];
	        this.startupError = source["startupError"];
	    }
	}

}

