-- https://blog.csdn.net/sinat_23931991/article/details/140350965

CREATE TABLE product (
                         id INTEGER NOT NULL PRIMARY KEY,
                         name VARCHAR(100),
                         attr1 JSONB,
                         attr2 JSON
);



INSERT INTO product
VALUES
    ( 1, '阿莫西林',
      '{
        "code": "100011",
        "price": 18,
        "uom": "盒",
        "specification": "10粒/盒",
        "price": 18.00
      }',
      '{
        "code": "100011",
        "price": 18,
        "uom": "盒",
        "specification": "10粒/盒",
        "price": 18.00
      }'
    );

INSERT INTO product
VALUES
    ( 2, '阿司匹林',
      '{
        "code": "100013",
        "price": 28,
        "uom": "盒",
        "specification": "20粒/盒",
        "price": 28.00
      }',
      '{
        "code": "100013",
        "price": 28,
        "uom": "盒",
        "specification": "20粒/盒",
        "price": 28.00
      }'
    );



INSERT INTO product
VALUES
    ( 3, '农夫山泉', jsonb_build_object ( 'code', '100015', 'price', 3 ),
      json_build_object ( 'code', '100015', 'price', 3 ) );