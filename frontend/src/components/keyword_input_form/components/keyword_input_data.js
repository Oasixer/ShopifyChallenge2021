export default class KeywordInputData {
  constructor(name, aliases, related, category, show_details){
    this.name = name;
    this.aliases = aliases;
    this.related = related;
    this.category = category;
    this.show_details = show_details;
  }
  static NewInst(){
    return new KeywordInputData('', [], [], false, true);
  }

  toJSON() {
    return JSON.stringify({name: this.name, aliases: this.aliases, related: this.related, category:this.category});
  }

  valid(all_keywords) {
    if (!this.name) {
      return {valid: false, message: 'missing keyword name'};
    }
    if (all_keywords.filter((i) => i.name === this.name).length > 1){
      return {valid: false, message: `keyword ${this.name} already exists`};
    }
    let invalid = [];
    this.related.forEach(i => {
      // TODO: check against master list of tags
      let found = false;
      let foundAlias = false;
      all_keywords.forEach(j => {
        if (j.name == this.name){
          found = true;
        }
      });
      if (!found){
        invalid.push(i.name);
      }
    });
    return {valid: invalid.length===0, message: `invalid params: ${JSON.stringify(invalid)}`};
  }
}
