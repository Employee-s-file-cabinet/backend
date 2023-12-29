BEGIN;

DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS organization_structure CASCADE;
DROP TABLE IF EXISTS positions CASCADE;
DROP TABLE IF EXISTS departments CASCADE;
DROP TABLE IF EXISTS benefits CASCADE;
DROP TABLE IF EXISTS benefit_uses CASCADE;
DROP TABLE IF EXISTS benefits_benefit_uses;
DROP TABLE IF EXISTS roles CASCADE;
DROP TABLE IF EXISTS authorizations;
DROP TABLE IF EXISTS educations;
DROP TABLE IF EXISTS trainings;
DROP TABLE IF EXISTS passports;
DROP TABLE IF EXISTS militaries;
DROP TABLE IF EXISTS vacations;
DROP TABLE IF EXISTS finances CASCADE;
DROP TABLE IF EXISTS work_types CASCADE;
DROP TABLE IF EXISTS contracts;
DROP TABLE IF EXISTS experiences;
DROP TABLE IF EXISTS indexations;

DROP TYPE IF EXISTS gender;
DROP TYPE IF EXISTS contract_type;

DROP FUNCTION IF EXISTS update_updated_at();

COMMIT;