<?php

namespace DataAccess;

use Symfony\Component\PropertyAccess\PropertyAccess;

class Collection
{

    protected $items;

    /**
     * @var \Symfony\Component\PropertyAccess\PropertyAccessor
     */
    protected $accessor;

    public function __construct()
    {
        $this->accessor = PropertyAccess::createPropertyAccessor();
    }

    /**
     * @param $data
     * @throws \Throwable
     */
    public function set($data)
    {
        $this->accessor->setValue($this->items, 'data', $data);
    }

    /**
     * @param $item
     * @return mixed
     */
    public function get()
    {
        return $this->accessor->getValue($this->items, 'data');
    }

}
