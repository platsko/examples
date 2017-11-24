<?php

namespace PetrLatsko\JsonBundle\Client;

use PetrLatsko\JsonBundle\Client\Exceptions\TransportRequestException;
use PetrLatsko\JsonBundle\Client\Interfaces\TransportInterface;

abstract class AbstractTransportProxy implements TransportInterface
{

    /**
     * @throws TransportRequestException
     */
    public function prepareErrors()
    {
        if ($this->hasErrors())
        {
            throw new TransportRequestException($this->getErrorMessage(), $this->getErrorCode());
        }
    }

}
