const fs = require('fs');
const path = require('path');
const { promisify } = require('util');

const readFile = promisify(fs.readFile);
const writeFile = promisify(fs.writeFile);

class ConfigParser {
  constructor(configPath) {
    this.configPath = path.resolve(configPath);
    this.cache = {};
  }

  async load() {
    try {
      const data = await readFile(this.configPath, 'utf8');
      this.cache = JSON.parse(data);
      return this.cache;
    } catch (error) {
      if (error.code === 'ENOENT') {
        await writeFile(this.configPath, JSON.stringify({}, null, 2));
        return {};
      }
      throw error;
    }
  }

  async save() {
    const data = JSON.stringify(this.cache, null, 2);
    await writeFile(this.configPath, data);
  }

  get(key) {
    return this.cache[key];
  }

  set(key, value) {
    this.cache[key] = value;
  }

  delete(key) {
    delete this.cache[key];
  }

  clear() {
    this.cache = {};
  }
}

module.exports = ConfigParser;