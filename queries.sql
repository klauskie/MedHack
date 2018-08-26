insert into medicamentos(nombre, descripcion) values ('Ibuprofeno', 'Ibuprofeno es una droga antiinflamatoria no esteroide (AINE o NSAID, por sus siglas en inglés). Esta medicina trabaja reduciendo sustancias en el cuerpo que causan dolor e inflamación.'),
('Diclofenac', 'Diclofenac es una droga antiinflamatoria no esteroide (AINE o NSAID, por sus siglas en inglés). Esta medicina trabaja reduciendo sustancias en el cuerpo que causan dolor e inflamación.'),
('Paracetamol', 'Paracetamol es un analgésico y un reductor de fiebre. El mecanismo exacto de acción de no se conoce.'),
('Dramamine', 'Antihistamínico que impide la propagación de impulsos emetógenos aferentes a nivel de núcleos vestibulares y anticolinérgico periférico que inhibe hipersecreción e hipermotilidad gástrica.'),
('Domperidona', 'La domperidona solo está indicada para el tratamiento de náuseas y vómitos.'),
('Bromfeniramina', 'Alivia ojos irritados, escurrimiento nasal y estornudos. Combate el catarro común.'),
('Dexametasona oftálmica','Potente corticosteroide con propiedades antialérgicas, antiexudativas y antiproliferativas inhibiendo la respuesta inflamatoria.');

insert into enfermedades(nombre) values
('Dolor de cabeza'),
('Mareo'),
('Dolor muscular'),
('Vomito'),
('Fiebre'),
('Dolor de Estomago'),
('Ojos irritados'),
('Escurrimiento nasal');

insert into enf_med(id_med, id_enf) values
(15,5),
(16,5),
(17,5),
(19,6),
(18,6),
(20,9),
(17,9),
(15,9),
(21,11),
(20,11),
(19,10);

insert into enf_med(id_med, id_enf) values
(17, 7),
(19, 8),
(18, 8);

SELECT enf_med.id_med, medicamentos.nombre FROM enfermedades join enf_med on enfermedades.id = enf_med.id_enf join medicamentos on medicamentos.id = enf_med.id_med where enfermedades.id = 5;


SELECT medicamentos.id, medicamentos.nombre, medicamentos.descripcion 
FROM enfermedades join enf_med on enfermedades.id = enf_med.id_enf 
join medicamentos on medicamentos.id = enf_med.id_med 
where enf_med.id_enf = 5;

SELECT medicamentos.id, medicamentos.nombre FROM enf_med join medicamentos on enf_med.id_med = medicamentos.id where enf_med.id_enf = 5;

SELECT medicamentos.id, medicamentos.nombre
FROM enf_med join medicamentos on enf_med.id_med = medicamentos.id
where enf_med.id_enf = 5;

insert into enfermedades(nombre) values
('Deshidratacion'),
('Embarazo');

insert into sintomas(nombre) values
('diarrea');

insert into enf_sint(id_enf, id_sint) values
(13,1),
(13,2),
(13,9);

insert into enf_sint(id_enf, id_sint) values
(14,1),
(14,2),
(14,4);

select * from enfermedades join enf_sint on enfermedades.id = enf_sint.id_enf where enfermedades.id=13;

select enfermedades.id, enfermedades.nombre from enfermedades join enf_sint on enfermedades.id = enf_sint.id_enf where enf_sint.id_sint = 1;

insert into enf_med(id_med, id_enf) values
(20, 12),
(15, 12);