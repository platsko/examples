<?php

namespace PetrLatsko\JsonBundle\Client\Interfaces;

use PetrLatsko\JsonBundle\Client\AbstractTransportProxy;

interface TransportClientInterface
{

    public function getTransport();

    public function setTransport(AbstractTransportProxy $transport);

    public function execRequest($url);

}
