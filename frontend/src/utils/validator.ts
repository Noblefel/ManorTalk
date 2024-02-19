export type FormErrors = Record<string, string[]>;

/** Validator implements a trivial form validations */
export class Validator {
  form: Record<string, any>;
  errors: FormErrors;

  constructor(form: Record<string, any>) {
    this.form = form;
    this.errors = {};
  }

  /** required will checks the given fields if they are empty */
  required(...fields: string[]) {
    for (const field of fields) {
      if (this.form[field] == "" || this.form[field] == undefined) {
        this.addError(field, "This field cannot be blank");
      }
    }
    return this;
  }

  /** email checks if the field value contains "@" symbol */
  email(field: string) {
    const email = this.form[field] as String;
    if (!email.includes("@")) {
      this.addError(field, "Is not a valid email");
    }
    return this;
  }

  /** strMinLength checks if the field value is less than the given number */
  strMinLength(field: string, n: number) {
    const str = this.form[field] as String;
    if (str?.length < n) {
      this.addError(field, `Must be atleast ${n} characters`);
    }
    return this;
  }

  /** strMaxLength checks if the field value is more than the given number */
  strMaxLength(field: string, n: number) {
    const str = this.form[field] as String;
    if (str?.length > n) {
      this.addError(field, `Must not exceeds ${n} characters`);
    }
    return this;
  }

  /** equal checks if a field has similar value to another*/
  equal(field: string, target: string) {
    if (this.form[field] != this.form[target]) {
      this.addError(field, `${field} does not match with ${target}`);
    }
    return this;
  }

  /** isValid will check if the form contains any error */
  isValid() {
    return Object.keys(this.errors).length == 0;
  }

  private addError(field: string, message: string) {
    this.errors[field] = [...(this.errors[field] || []), message];
  }
}
