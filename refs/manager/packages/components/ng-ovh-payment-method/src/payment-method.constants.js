export const DEFAULT_OPTIONS = {
  onlyValid: false,
  transform: true, // transform legacy payment methods to payment methods
};

export const DEFAULT_TYPE_OPTIONS = {
  onlyRegisterable: true,
  transform: true, // transform legacy payment method types to payment method types
};

export const IBAN_BIC_RULES = {
  IBAN_FORMAT: {
    FR: [4, 4, 4, 4, 4, 4, 3],
    DE: [4, 4, 4, 4, 4, 2],
    MC: [4, 4, 4, 4, 4, 4, 3],
  },
  IBAN_VALIDATION_MODULO: 97,
  COUNTRY_BASE_REGEXP: {
    AT: new RegExp(/\d{16}/),
    BE: new RegExp(/\d{12}/),
    BG: new RegExp(/\w{4}\d{6}[0-9A-Z]{8}/),
    CH: new RegExp(/\d{5}[0-9A-Z]{12}/),
    CY: new RegExp(/\d{8}[0-9A-Z]{16}/),
    CZ: new RegExp(/\d{20}/),
    DE: new RegExp(/\d{18}/),
    DK: new RegExp(/\d{14}/),
    EE: new RegExp(/\d{16}/),
    ES: new RegExp(/\d{20}/),
    FI: new RegExp(/\d{14}/),
    FR: new RegExp(/\d{10}\w{11}\d{2}/),
    GR: new RegExp(/\d{7}[0-9A-Z]{16}/),
    HU: new RegExp(/\d{24}/),
    IE: new RegExp(/\w{4}\d{14}/),
    IS: new RegExp(/\d{22}/),
    IT: new RegExp(/\w{1}\d{10}[0-9A-Z]{12}/),
    LI: new RegExp(/\d{5}[0-9A-Z]{12}/),
    LT: new RegExp(/\d{16}/),
    LU: new RegExp(/\d{3}[0-9A-Z]{13}/),
    LV: new RegExp(/\w{4}[0-9A-Z]{13}/),
    MC: new RegExp(/\d{10}[0-9A-Z]{11}\d{2}/),
    MT: new RegExp(/\w{4}\d{5}[0-9A-Z]{18}/),
    NL: new RegExp(/\w{4}\d{10}/),
    NO: new RegExp(/\d{9}/),
    PL: new RegExp(/\d{8}[0-9A-Z]{16}/),
    PT: new RegExp(/\d{21}/),
    RO: new RegExp(/\w{4}[0-9A-Z]{16}/),
    SE: new RegExp(/\d{20}/),
    SI: new RegExp(/\d{15}/),
    SK: new RegExp(/\d{20}/),
    UK: new RegExp(/\w{4}\d{14}/),
  },
  BIC_REGEXP: /^([A-Z]{4})([A-Z]{2})(\w{2})(\w{3})?$/g,
  IBAN_REGEXP: /^([A-Z]{2})(\d{2})(.*)$/g,
};

export const PAYMENT_METHOD_STATUS_ENUM = {
  CANCELED: 'CANCELED',
  CANCELING: 'CANCELING',
  CREATED: 'CREATED',
  ERROR: 'ERROR',
  EXPIRED: 'EXPIRED',
  CREATING: 'CREATING',
  MAINTENANCE: 'MAINTENANCE',
  PAUSED: 'PAUSED',
  VALID: 'VALID',
};

export const PAYMENT_METHOD_TYPE_ENUM = {
  BANK_ACCOUNT: 'BANK_ACCOUNT',
  CREDIT_CARD: 'CREDIT_CARD',
  DEFERRED_PAYMENT_ACCOUNT: 'DEFERRED_PAYMENT_ACCOUNT',
  PAYPAL: 'PAYPAL',
};

export const TYPE_INTEGRATION_ENUM = {
  NONE: 'NONE',
  REDIRECT: 'REDIRECT',
  IFRAME_VANTIV: 'IFRAME_VANTIV',
  IN_CONTEXT: 'IN_CONTEXT',
};

export default {
  DEFAULT_OPTIONS,
  DEFAULT_TYPE_OPTIONS,
  IBAN_BIC_RULES,
  PAYMENT_METHOD_STATUS_ENUM,
  PAYMENT_METHOD_TYPE_ENUM,
  TYPE_INTEGRATION_ENUM,
};
