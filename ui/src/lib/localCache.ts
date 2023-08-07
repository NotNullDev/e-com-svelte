class LocalCache {
  private static instance: LocalCache;
  private cache: Map<string, any>;

  private constructor() {
    this.cache = new Map<string, string>();
  }

  public static getInstance(): LocalCache {
    if (!LocalCache.instance) {
      LocalCache.instance = new LocalCache();
    }

    return LocalCache.instance;
  }

  public set(key: string, value: any): void {
    this.cache.set(key, value);
  }

  public get(key: string): string | undefined {
    return this.cache.get(key);
  }
}
