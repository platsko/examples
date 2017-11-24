<?php

namespace DataAccess;

use ArrayIterator;
use IteratorAggregate;

class Data implements IteratorAggregate
{

    private $data = [];

    /**
     * @param array $data
     */
    public function setData(array $data)
    {
        $this->data = $data;
    }

    /**
     * @return ArrayIterator|\Traversable
     */
    public function getData()
    {
        return $this->getIterator();
    }

    /**
     * @return ArrayIterator|\Traversable
     */
    public function getIterator()
    {
        return new ArrayIterator($this->data);
    }
}
