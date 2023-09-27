export namespace youtube {
	
	export class Format {
	    itag: number;
	    url: string;
	    mimeType: string;
	    quality: string;
	    signatureCipher: string;
	    bitrate: number;
	    fps: number;
	    width: number;
	    height: number;
	    lastModified: string;
	    contentLength: number;
	    qualityLabel: string;
	    projectionType: string;
	    averageBitrate: number;
	    audioQuality: string;
	    approxDurationMs: string;
	    audioSampleRate: string;
	    audioChannels: number;
	    // Go type: struct { Start string "json:\"start\""; End string "json:\"end\"" }
	    initRange?: any;
	    // Go type: struct { Start string "json:\"start\""; End string "json:\"end\"" }
	    indexRange?: any;
	
	    static createFrom(source: any = {}) {
	        return new Format(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.itag = source["itag"];
	        this.url = source["url"];
	        this.mimeType = source["mimeType"];
	        this.quality = source["quality"];
	        this.signatureCipher = source["signatureCipher"];
	        this.bitrate = source["bitrate"];
	        this.fps = source["fps"];
	        this.width = source["width"];
	        this.height = source["height"];
	        this.lastModified = source["lastModified"];
	        this.contentLength = source["contentLength"];
	        this.qualityLabel = source["qualityLabel"];
	        this.projectionType = source["projectionType"];
	        this.averageBitrate = source["averageBitrate"];
	        this.audioQuality = source["audioQuality"];
	        this.approxDurationMs = source["approxDurationMs"];
	        this.audioSampleRate = source["audioSampleRate"];
	        this.audioChannels = source["audioChannels"];
	        this.initRange = this.convertValues(source["initRange"], Object);
	        this.indexRange = this.convertValues(source["indexRange"], Object);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

