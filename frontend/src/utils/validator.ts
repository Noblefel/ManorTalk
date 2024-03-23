export type FormErrors = Record<string, string[]>;

/** Validator implements a trivial form validations */
export class Validator {
  form: Record<string, any>;
  errors: FormErrors;

  constructor(form: Record<string, any>) {
    this.form = form;
    this.errors = {};
  }

  required(...fields: string[]) {
    for (const field of fields) {
      if (this.form[field] == "" || this.form[field] == undefined) {
        this.addError(field, "This field cannot be blank");
      }
    }
    return this;
  }

  email(field: string) {
    const email = this.form[field] as String;
    if (!email.includes("@")) {
      this.addError(field, "Is not a valid email");
    }
    return this;
  }

  strMinLength(field: string, n: number) {
    const str = this.form[field] as String;
    if (str?.length < n) {
      this.addError(field, `Must be atleast ${n} characters`);
    }
    return this;
  }

  strMaxLength(field: string, n: number) {
    const str = this.form[field] as String;
    if (str?.length > n) {
      this.addError(field, `Must not exceeds ${n} characters`);
    }
    return this;
  }

  equal(field: string, target: string) {
    if (this.form[field] != this.form[target]) {
      this.addError(field, `${field} does not match with ${target}`);
    }
    return this;
  }

  isValid() {
    return Object.keys(this.errors).length == 0;
  }

  private addError(field: string, message: string) {
    this.errors[field] = [...(this.errors[field] || []), message];
  }
}
